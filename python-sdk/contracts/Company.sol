pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

import "./Table.sol";

contract Company {
  event CreateResult(int count);
  event InsertResult(int count);
  event UpdateResult(int count);
  event RemoveResult(int count);

  function createCompany() public returns (int) {
    TableFactory tf = TableFactory(0x1001);
    int count = tf.createTable("Company", "key", "cid,name,trustworthy,funds");
    emit CreateResult(count);
    return count;
  }

  function createAccount() public returns (int) {
    TableFactory tf = TableFactory(0x1001);
    int count = tf.createTable("Account", "key", "aid,A,B,Money");
    emit CreateResult(count);
    return count;
  }



  //select records
  function selectCompany() public returns(int[], string[], int[], int[]){

    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Company");

    Condition condition = table.newCondition();
    Entries entries = table.select("key", condition);

    int[] memory company_id_list = new int[](uint256(entries.size()));
    string[] memory company_name_list = new string[](uint256(entries.size()));
    int[] memory company_trustworthy_list = new int[](uint256(entries.size()));
    int[] memory company_funds_list = new int[](uint256(entries.size()));

    for(int i = 0; i < entries.size(); ++i) {
      Entry entry = entries.get(i);
      company_id_list[uint256(i)] = entry.getInt("cid");
      company_name_list[uint256(i)] = entry.getString("name");
      company_trustworthy_list[uint256(i)] = entry.getInt("trustworthy");
      company_funds_list[uint256(i)] = entry.getInt("funds");
    }
    return (company_id_list, company_name_list, company_trustworthy_list, company_funds_list);
  }

  //insert records
  function insertCompany(int id, string name, int trustworthy, int funds) public returns(int) {
    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Company");

    Entry entry = table.newEntry();
    entry.set("cid", id);
    entry.set("name", name);
    entry.set("trustworthy", trustworthy);
    entry.set("funds", funds);

    int count = table.insert("key", entry);

    emit InsertResult(count);
    return count;
  }
    //update records
  function updateCompany(int id, string name, int trustworthy, int funds) public returns(int) {
    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Company");

    Entry entry = table.newEntry();
    entry.set("cid", id);
    entry.set("name", name);
    entry.set("funds", funds);
    entry.set("trustworthy", trustworthy);

    Condition condition = table.newCondition();
    condition.EQ("cid", id);

    int count = table.update("key", entry, condition);

    emit UpdateResult(count);
    return count;
  }
    //remove records
  function removeCompany(int id) public returns(int){
    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Company");

    Condition condition = table.newCondition();
    condition.EQ("cid", id);

    int count = table.remove("key", condition);

    emit RemoveResult(count);
    return count;
  }

  function selectAccount() public returns(int[], int[], int[], int[]) {
    Account account = new Account();
    return account.select();
  }

  function insertAccount(int id, int A, int B, int money) public returns(int) {
    Account account = new Account();
    return account.insert(id, A, B, money);
  }

  function updateAccount(int id, int A, int B, int money) public returns(int) {
    Account account = new Account();
    return account.update(id, A, B, money);
  }

  function removeAccount(int id) public returns(int) {
    Account account = new Account();
    return account.remove(id);
  }
}



contract Account {
  event InsertResult(int count);
  event UpdateResult(int count);
  event RemoveResult(int count);
  //select records
  function select() public returns(int[], int[], int[], int[]){

    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Account");

    Condition condition = table.newCondition();

    Entries entries = table.select("key", condition);

    int[] memory id_list = new int[](uint256(entries.size()));
    int[] memory A_list = new int[](uint256(entries.size()));
    int[] memory B_list = new int[](uint256(entries.size()));
    int[] memory money_list = new int[](uint256(entries.size()));

    for(int i = 0; i < entries.size(); i++) {
      Entry entry = entries.get(i);
      id_list[uint256(i)] = entry.getInt("aid");
      A_list[uint256(i)] = entry.getInt("A");
      B_list[uint256(i)] = entry.getInt("B");
      money_list[uint256(i)] = entry.getInt("Money");
    }
    return (id_list, A_list, B_list, money_list);
  }

  //insert records
  function insert(int id, int A, int B, int money) public returns(int) {
    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Account");

    Entry entry = table.newEntry();
    entry.set("aid", id);
    entry.set("A", A);
    entry.set("B", B);
    entry.set("Money", money);

    int count = table.insert("key", entry);

    emit InsertResult(count);
    return count;
  }
    //update records
  function update(int id, int A, int B, int money) public returns(int) {
    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Account");

    Entry entry = table.newEntry();
    entry.set("aid", id);
    entry.set("Money", money);
    entry.set("A", A);
    entry.set("B", B);

    Condition condition = table.newCondition();
    condition.EQ("aid", id);

    int count = table.update("key", entry, condition);

    emit UpdateResult(count);
    return count;
  }
    //remove records
  function remove(int id) public returns(int){
    TableFactory tf = TableFactory(0x1001);
    Table table = tf.openTable("Account");

    Condition condition = table.newCondition();
    condition.EQ("aid", id);

    int count = table.remove("key", condition);

    emit RemoveResult(count);
    return count;
  }
}

contract AccountFactory {
  function openAccount() public returns (Account); //open table
}
