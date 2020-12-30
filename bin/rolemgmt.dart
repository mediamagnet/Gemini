import 'package:mongo_dart/mongo_dart.dart';
import 'package:nyxx/nyxx.dart';

Future AddRole(Snowflake user, String role) async {
  var db = Db('mongodb://localhost:27017/gemini');
  await db.open();

  var coll = db.collection('role');
  await coll.insertAll([
    {'user': user, 'role': role}
  ]);
}