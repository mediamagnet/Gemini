import 'dart:io' show Platform;
import 'package:toml/loader/fs.dart';

String get dartVersion {
  final platformVersion = Platform.version;
  return platformVersion.split('(').first;
}

String helpCommandGen(String commandName, String description, { String? additionalInfo }) {
  FilesystemConfigLoader.use();
  final buffer = StringBuffer();
  var cfg;

  buffer.write('**${cfg['Bot']['Prefix']}${commandName}**');

  if (additionalInfo != null) {
    buffer.write(' `$additionalInfo`');
  }

  buffer.write(' - $description.\n');

  return buffer.toString();
}