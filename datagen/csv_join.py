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
toi = pd.read_csv(f"../raw_data/{teamname}_skaters_toi_{year}.csv")
goalies = pd.read_csv(f"../raw_data/{teamname}_goalies_{year}.csv")

# Modify Regular stats
regular.rename(columns = {  'EV':'EV_G', 'PP':'PP_G', 'SH':'SH_G',
                            'EV.1':'EV_A', 'PP.1':'PP_A', 'SH.1':'SH_A'}, inplace = True)

# Modify TOI stats
toi = toi.loc[:, ~toi.columns.str.contains('^Unnamed')]
toi.rename(columns = {  'TOI':'EV_TOI', 'CF% Rel':'EV_CF% Rel', 'GF/60':'EV_GF/60', 'GA/60':'EV_GA/60',
                        'TOI.1':'PP_TOI', 'CF% Rel.1':'PP_CF% Rel', 'GF/60.1':'PP_GF/60', 'GA/60.1':'PP_GA/60',
                        'TOI.2':'SH_TOI', 'CF% Rel.2':'SH_CF% Rel', 'GF/60.2':'SH_GF/60', 'GA/60.2':'SH_GA/60'
}, inplace = True)

# Drop duplicate columns
regular = regular.drop(['Rk'], axis=1)
advanced = advanced.drop(['Rk', 'Age', 'Pos', 'GP', '-9999'], axis=1)
toi = toi.drop(['Rk', 'Pos', 'GP', '-9999'], axis=1)
goalies = goalies.drop(['Rk', 'Age', 'GP', '-9999'], axis=1)

# Join tables
stats = regular.merge(advanced, how='outer', on='Player')
stats = stats.merge(toi, how='outer', on='Player')
stats = stats.merge(goalies, how='outer', on='Player')

# Write to file
stats.to_csv(f"../data/{teamname}_players_{year}.csv", encoding='utf-8', index=False)