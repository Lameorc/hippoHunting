import requests
from bs4 import BeautifulSoup
import re
from icalendar import Calendar, Event

result = requests.get("http://rishada.cz/turnaje")
if result.status_code == 200:
    soup = BeautifulSoup(result.content, "html.parser")
    """All booster drafts on page"""
    for elem in soup(text=re.compile(r'Booster Draft')):
        text = elem.parent.text
        regex = re.compile(r"\d{2}\.\d{2}\.\d{4}\D*- Booster Draft")
        events = sorted(set(regex.findall(text)))
        for event in events:
            regex = re.compile(r"(?P<day>\d{2})\.(?P<month>\d{2})\.(?P<year>\d{4})\D* - (?P<name>Booster Draft)")
            match = regex.search(event)
            cal = Calendar()
            calEvent = Event
            calEvent.add('dtstart', datetime(match.group(''))
            