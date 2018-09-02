import base64
import json
import logging
from flask import Flask


log = logging.getLogger("jira-board-merger")

app = Flask(__name__)
app.secret_key = "test"


@app.before_first_request
def initialize():
    logger = logging.getLogger("jira-board-merger")
    logger.setLevel(logging.DEBUG)
    ch = logging.StreamHandler()
    ch.setLevel(logging.DEBUG)
    formatter = logging.Formatter(
        """%(levelname)s in %(module)s [%(pathname)s:%(lineno)d]:\n%(message)s"""
    )
    ch.setFormatter(formatter)
    logger.addHandler(ch)


test_data = {
    "rapidViewId":70,
    "statistics":{"fieldConfigured":False,"typeId":"none","id":"","name":"None"},
    "columnsData":{
        "rapidViewId":70,
        "columns":[
            {
                "id":532,
                "name":"ToDo",
                "statusIds":["1","4"],
                "isKanPlanColumn":False
            },
            {
                "id":596,
                "name":"Waiting for",
                "statusIds":["2"],
                "isKanPlanColumn":False
            },
            {
                "id":533,
                "name":"In Progress",
                "statusIds":["3"],
                "isKanPlanColumn":False
            },
            {
            "id":586,
            "name":"Review",
            "statusIds":["7"],
            "isKanPlanColumn":False
            },
            {
                "id":534,
                "name":"Done",
                "statusIds":["6","5"],
                "isKanPlanColumn":False
            }
        ]
    },
    "swimlanesData":{"rapidViewId":70,"swimlaneStrategy":"assigneeUnassignedFirst"},
    "issuesData":{
        "rapidViewId":70,
        "activeFilters":[{"id":500},{"id":347}],
        "issues":[
            {
                "id":57996,
                "key":"PRJ-30700",
                "hidden":False,
                "typeName":"Task",
                "typeId":"3",
                "summary":"A debug task",
                "typeUrl":"https://jira.org/jira/task.gif",
                "done":False,
                "assignee":"user@company.com",
                "assigneeName":"User Name",
                "avatarUrl":"https://jira/jira/secure/useravatar?ownerId=user@company.com&avatarId=11245",
                "hasCustomUserAvatar":True,
                "color":"#f50a0a",
                "epic":"PRJ-18001",
                "epicField":{
                    "id":"customfield_2",
                    "label":"Epic",
                    "editable":False,
                    "renderer":"epiclink",
                    "epicKey":"PRJ-18001",
                    "epicColor":"ghx-label-7",
                    "text":"My Epic",
                    "canRemoveEpic":True
                },
                "estimateStatistic":{"statFieldId":"customfield_10232","statFieldValue":{}},
                "trackingStatistic":{"statFieldId":"timeestimate","statFieldValue":{"value":7200.0,"text":"2h"}},
                "statusId":"3",
                "statusName":"In Progress",
                "statusUrl":"https://jira/jira/images/icons/statuses/inprogress.png",
                "status":{
                    "id":"3",
                    "name":"In Progress",
                    "description":"This issue is being actively worked on at the moment by the assignee.",
                    "iconUrl":"https://jira/jira/images/icons/statuses/inprogress.png",
                    "statusCategory":{"id":"4","key":"indeterminate","colorName":"yellow"}
                },
                "fixVersions":[],
                "projectId":1,
                "linkedPagesCount":0,
                "extraFields":[
                    {"id":"customfield_10160","label":"Milestone","editable":False,"renderer":"html","html":"alpha"},
                    {"id":"timeestimate","label":"Remaining Estimate","editable":False,"renderer":"html","html":"2 hours"}
                ]
            },
        ],
        "projects":[{"id":1}],
        "missingParents":[],
        "canRelease":False,
        "hasBulkChangePermission":True
    },
    "orderData":{
        "rapidViewId":70,
        "rankable":True,
        "rankCustomFieldId":10440,
        "canRankPerProject":[{"projectId":1,"projectKey":"PRJ","canRank":True}]
    },
    "sprintsData":{
        "rapidViewId":70,
        "sprints":[
            {
                "id":72,
                "sequence":71,
                "name":"Sprint1",
                "state":"ACTIVE",
                "linkedPagesCount":0,
                "startDate":"19/Aug/10 6:00 AM",
                "endDate":"22/Sep/10 6:00 AM",
                "completeDate":"None",
                "canUpdateSprint":False,
                "remoteLinks":[],
                "daysRemaining":5
            }
        ],
        "canManageSprints":False
    },
    "etagData":{
        "rapidViewId":70,"issueCount":217,"lastUpdated":1535366717000,"quickFilters":"[500, 347]","sprints":"[]","etag":"70,1535366717000,[500, 347],[],217"
    }
}


@app.route('/allData', methods=['GET'])
def index():
    return json.dumps(test_data)


if __name__ == "__main__":
    app.run()
