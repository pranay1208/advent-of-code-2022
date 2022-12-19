import re

with open("./input.txt") as file:
    inp = file.read().strip().split('\n')

def getKey(valve: str, time: int) -> str:
    return f"{valve}|{time}"

# 1. Parse input
pattern = re.compile(
    r"Valve (?P<valve>.*) has flow rate=(?P<rate>\d+); tunnels? leads? to valves? (?P<routes>.*)")

routeMap = {}
rateMap = {}
maxMap = {}

for line in inp:
    match = pattern.match(line)
    valve: str = match.group('valve')
    rate = int(match.group('rate'))
    routes = match.group('routes').split(', ')
    routeMap[valve] = routes
    rateMap[valve] = rate
    # 2. Create map and pre-fill values for 0 and 1
    maxMap[getKey(valve, 0)] = 0
    maxMap[getKey(valve, 1)] = rate


# 3. For any t >= 2 and t <= 30, compute by seeing max of directly moving to other room versus switching on valve and moving to other room
for t in range(2,31):
    for valve in rateMap:
        # max value if I open this valve and then go somewhere else
        val1 = rateMap[valve] * (t-1) +  max(list(
            map(lambda v: maxMap[getKey(v, t-2)], routeMap[valve])
            ))
        # max value if I go somewhere else directly
        val2 = max(list(
            map(lambda v: maxMap[getKey(v, t-1)], routeMap[valve])
        ))

        maxMap[getKey(valve, t)] = max(val1, val2)

# 4. Answer is valve AA at t = 30
print(maxMap[getKey('AA', 30)])
