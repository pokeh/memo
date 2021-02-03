## Unique key

```
ALTER TABLE `table_name`
   ADD UNIQUE KEY `col1` (`col1`, `col2`, `col3`);
```

or `KEY` can be omitted, like:

```
ALTER TABLE `table_name`
  ADD UNIQUE `col1` (`col1`, `col2`, `col3`);
```

- General terminology for "keys that ensure uniqueness"
- Key == index

## Primary key

```
ALTER TABLE `table_name`
  ADD PRIMARY KEY (`col1`);
```

- A special type of unique key:
  - Each table can have only one primary key
  - The key has to be a single field

## Unique index

```
ALTER TABLE `table_name`
  ADD UNIQUE INDEX `col1` (`col1`, `col2`, `col3`);
```

- Unlike primary keys:
  - A table can have more than one unique index
  - A unique index can use multiple fields

## Unique constraint

```
ALTER TABLE table_name
  ADD CONSTRAINT constraint_name UNIQUE (col1, col2);
```

- This creates a corresponding unique index
- Constraint: something that restrict the values that are stored in a table
  - There are column level constraints and table level constraints
  - Other constraints include NOT NULL, DEFAULT, AUTO_INCREMENT, FOREIGN_KEY, etc.
