import 'package:nyxx/nyxx.dart';
import 'package:nyxx_commander/commander.dart'
    show CommandContext, CommandGroup, Commander;
import 'package:toml/loader/fs.dart';
import 'dart:async';
import 'dart:math' show Random;
import 'dart:io' show Process, ProcessInfo, pid, sleep;
import 'utils.dart' as utils;
import 'rolecmd.dart' as role;


var ownerID;
var secOwner;
var prefix;

Future main(List<String> arguments) async {
  FilesystemConfigLoader.use();
  var cfg;
  try {
    cfg = await loadConfig('config.toml');
    print(cfg['Bot']['ID']);
    ownerID = cfg['Owner']['ID'][0];
    secOwner = cfg['Owner']['ID'][1];
    prefix = cfg['Bot']['Prefix'];

    final bot = Nyxx(cfg['Bot']['Token']!);

    bot.onMessageReceived.listen((MessageReceivedEvent e){
      if (e.message.content.contains('<@!783065070682243082>')) {
        print(e.message.content);
        e.message.createReaction(UnicodeEmoji('❤'));
      } else if (e.message.content.contains('Good bot')) {
        print(e.message.content);
        e.message.createReaction(UnicodeEmoji('❤'));
      } else if (e.message.content.contains('Bad bot')) {
        print(e.message.content);
        e.message.createReaction(UnicodeEmoji('🗡'));
      }
    });
    bot.onReady.listen((ReadyEvent e) {
      print('Connected to discord.');
      bot.setPresence(PresenceBuilder.of(game: Activity.of(
          'Gemini v1', type: ActivityType.streaming,
          url: 'https://github.com/mediamagnet/gemini')));
    });

    Commander(bot, prefix: cfg['Bot']['Prefix'])
      ..registerCommandGroup(CommandGroup(beforeHandler: checkForAdmin)
        ..registerSubCommand('shutdown', shutdownCommand))
      ..registerCommand('ping', pingCommand)
      ..registerCommand('dinfo', dinfoCommand)
      ..registerCommand('help', helpCommand)
      ..registerCommandGroup(CommandGroup(name: 'role')
        ..registerSubCommand('add', role.addCommand));

  } catch (e) {
    print(e);
  }
  return cfg;
}

Future<void> dinfoCommand(CommandContext ctx, String content) async {
  final color = DiscordColor.fromRgb(
      Random().nextInt(255), Random().nextInt(255), Random().nextInt(255));

  final embed = EmbedBuilder()
    ..addAuthor((author) {
      author.name = ctx.client.self.tag;
      author.iconUrl = ctx.client.self.avatarURL();
      author.url = 'https://github.com/mediamagnet/gemini-dart';
    })
    ..addFooter((footer) {
      footer.text =
      'Gemini | Shard [${ctx.shardId + 1}] of [${ctx.client.shards}] | ${utils.dartVersion}';
    })
    ..color = color
    ..addField(
        name: 'Uptime', content: ctx.client.uptime.inMinutes, inline: true)
    ..addField(
        name: 'DartVM memory usage',
        content:
        '${(ProcessInfo.currentRss / 1024 / 1024).toStringAsFixed(2)} MB',
        inline: true)
    ..addField(
        name: 'Created at', content: ctx.client.app.createdAt, inline: true)
    ..addField(
        name: 'Guild count', content: ctx.client.guilds.count, inline: true)
    ..addField(
        name: 'Users count', content: ctx.client.users.count, inline: true)
    ..addField(
        name: 'Channels count',
        content: ctx.client.channels.count,
        inline: true)
    ..addField(
        name: 'Users in voice',
        content: ctx.client.guilds.values
            .map((g) => g.voiceStates.count)
            .reduce((f, s) => f + s),
        inline: true)
    ..addField(name: 'Shard count', content: ctx.client.shards, inline: true)
    ..addField(
        name: 'Cached messages',
        content: ctx.client.channels
            .find((item) => item is MessageChannel)
            .cast<MessageChannel>()
            .map((e) => e.messages.count)
            .fold(0, (first, second) => (first as int) + second),
        inline: true);

  await ctx.message.delete();
  await ctx.reply(embed: embed);
}

Future<void> shutdownCommand(CommandContext ctx, String content) async {
  await ctx.message.delete();
  await ctx.reply(content: 'I...');
  sleep(Duration(seconds: 5));
  await ctx.reply(content: 'I guess I can... bye.');
  Process.killPid(pid);
}

Future<void> helpCommand(CommandContext ctx, String content) async {
  final color = DiscordColor.fromRgb(
      Random().nextInt(255), Random().nextInt(255), Random().nextInt(255));

  // Write zero-width character to skip first line where nick is
  final embed = EmbedBuilder()
    ..addAuthor((author) {
      author.name = ctx.client.self.tag;
      author.iconUrl = ctx.client.self.avatarURL();
      author.url = 'https://github.com/mediamagnet/gemini-dart';
    })
    ..addFooter((footer) {
      footer.text =
      'Gemini | Shard [${ctx.shardId + 1}] of [${ctx.client.shards}] | ${utils.dartVersion}';
    })
    ..color = color
    ..addField(name: 'Ping', content: 'Sends a ping response.')
    ..addField(name: 'Help', content: "You're reading it.")
    ..addField(name: 'dinfo', content: 'Sends info about the dart side');

  await ctx.message.delete();
  await ctx.reply(embed: embed);
}

Future<void> pingCommand(CommandContext ctx, String content) async {
  await ctx.message.delete();
  final random = Random();
  final color = DiscordColor.fromRgb(
      random.nextInt(255), random.nextInt(255), random.nextInt(255));
  final gatewayDelayInMilis = ctx.client.shardManager.shards
      .firstWhere((element) => element.id == ctx.shardId)
      .gatewayLatency
      .inMilliseconds;
  final stopwatch = Stopwatch()..start();

  final embed = EmbedBuilder()
    ..color = color
    ..addField(
        name: 'Gateway latency',
        content: '$gatewayDelayInMilis ms',
        inline: true)
    ..addField(
        name: 'Message roundup time', content: 'Pending...', inline: true);

  final message = await ctx.reply(embed: embed);

  embed
    ..replaceField(
        name: 'Message roundup time',

        content: '${stopwatch.elapsedMilliseconds} ms',
        inline: true);

  await message.edit(embed: embed);
}

Future<bool> checkForAdmin(CommandContext context) async {
  if (ownerID != null) {
    return context.author!.id == ownerID;
  } else if (secOwner != null) {
    return context.author!.id == secOwner;
  }
  return false;
}

