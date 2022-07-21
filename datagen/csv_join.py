import pandas as pd
import sys

# Check args1
if len(sys.argv) != 3:
    print("Incorrect number of arguments. Usage: `csv_join <teamname> <year>")
    exit()

# Read args
[teamname, year] = sys.argv[1:]

# Read from csv
regular = pd.read_csv(f"../raw_data/{teamname}_skaters_{year}.csv")
advanced = pd.read_csv(f"../raw_data/{teamname}_skaters_adv_{year}.csv")
goalies = pd.read_csv(f"../raw_data/{teamname}_goalies_{year}.csv")

# Drop duplicate columns
regular = regular.drop(['Rk'], axis=1)
advanced = advanced.drop(['Rk', 'Age', 'Pos', 'GP', '-9999'], axis=1)
goalies = goalies.drop(['Rk', 'Age', 'GP', '-9999'], axis=1)

# Join tables
stats = regular.merge(advanced, how='outer', on='Player')
stats = stats.merge(goalies, how='outer', on='Player')

# Write to file
stats.to_csv(f"../data/{teamname}_skaters_{year}.csv", encoding='utf-8', index=False)