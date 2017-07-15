// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	"io"
)

// Broker execute SQL statements in the data store.
// It marshals/un-marshals go structures.
type Broker interface {
	// Put puts single value (inBinding) into the data store
	// Example usage:
	//
	//    err = db.Put("ID='James Bond'", &User{"James Bond", "James", "Bond"})
	//
	Put(inBinding interface{}, /* TODO opts ...PutOption*/) error

	// NewTxn creates a transaction / batch
	NewTxn() Txn

	// GetValue retrieves one item based on the query. If the item exists it is un-marshaled into the outBinding.
	// Example usage:
	//
	//    query := sql.SelectFrom(UserTable) + sql.Where(sql.FieldEq(&UserTable.ID, UserTable, "James Bond"))
	//    user := &User{}
	//    found, err := db.GetValue(query, user)
	//
	GetValue(query string, outBinding interface{}) (found bool, err error)

	// ListValues returns an iterator that enables to traverse all items returned by the query
	// Example usage:
	//
	//    query := sql.SelectFrom(UserTable) + sql.Where(sql.FieldEq(&UserTable.LastName, UserTable, "Bond"))
	//    iterator := db.ListValues(query)
	//    users := &[]User{}
	//    err := sql.SliceIt(users, iterator)
	//
	ListValues(query string) ValIterator

	// Delete removes data that from the data store
	// Example usage:
	//
	//    err := db.Delete("from User wher eID='James Bond'")
	//
	Delete(fromWhere string) error

	// Executes the SQL statement (can be used for example for create "table/type" if not exits...)
	// Example usage:
	//
	//  	 err := db.Exec("CREATE INDEX IF NOT EXISTS...")
	Exec(statement string) error
}

// ValIterator is an iterator returned by ListValues call.
type ValIterator interface {
	// GetNext retrieves the current "row" from query result. GetValue is un-marshaled into the provided argument.
	// The stop=true will be returned if there is no more record or if error occurred (to get the error call Close())
	// Whe the stop=true is returned the outBinding was not updated.
	GetNext(outBinding interface{}) (stop bool)

	// Closer is used to retrieve error (if occurred) & release the cursor
	io.Closer
}

// Txn allows to group operations into the transaction or batch (depending on a particular data store).
// Transaction executes usually multiple operations in a more efficient way in contrast to executing them one by one.
type Txn interface {
	// Put adds put operation into the transaction
	Put(where string, data interface{}) Txn
	// Delete adds delete operation, which removes value identified by the key, into the transaction
	Delete(where string) Txn
	// Commit tries to commit the transaction.
	Commit() error
}
