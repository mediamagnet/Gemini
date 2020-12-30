import 'package:nyxx/nyxx.dart';
import 'package:nyxx_commander/commander.dart';
import 'rolemgmt.dart' as role;

Future<void> addCommand(CommandContext ctx, String content) async {

  var roleID = RegExp('<@&(?<id>\d{4,})>').firstMatch(content)!.namedGroup('id');
  await role.AddRole(ctx.author!.id, roleID.toString());
}
