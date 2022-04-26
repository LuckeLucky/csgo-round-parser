import json
import os
import glob

path = './jsons'
for filename in glob.glob(os.path.join(path, '*.json')):
    with open(os.path.join(os.getcwd(), filename), 'r') as f: 
        data = json.load(f)
        players = data["Players"]
        for player in players:
            print('[['+player['Page']+']]')
            print('<pre class="selectall">')
            print('{{Crosshair table')
            print('|date='+data['Date'])
            print('|ref=<ref>{{cite_web|date='+data['Date']+'|url='+data['Link']+'|title='+data['Title']+'|publisher=[[HLTV]]}}</ref>')
            print('|share_code='+player['CrosshairCode'])
            print('}}')
            print('</pre>')
