select t.*, eau.name as author, eas.name as assignee
from tasks t left join
     employees eau
     on eau.id = t.author_id left join
     employees eas
     on eas.id = t.assignee_id;


     SELECT m.ID,
       m.SCORE_HOME,
       m.SCORE_AWAY,
       t1.NAME as name_home_team,
       t2.NAME as name_home_team
FROM MATCHES AS m
JOIN teams AS t1 ON t1.id = m.home_team_id
JOIN teams AS t2 ON t2.id = m.away_team_id
WHERE m.id = 1

