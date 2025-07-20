SELECT
    p.id AS person_id,
    ni.id AS named_identifier_id,
    ni.first_name,
    ni.surname,
    ni.nickname
FROM
    person AS p
JOIN
    person_name_identifier AS pni ON p.id = pni.person_id
JOIN
    named_identifier AS ni ON pni.named_identifier_id = ni.id;
