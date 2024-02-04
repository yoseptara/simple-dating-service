DROP INDEX IF EXISTS idx_swipes_user_id_swipe_date;

ALTER TABLE swipes
DROP CONSTRAINT IF EXISTS fk_swipes_user_id;

DROP TABLE IF EXISTS swipes;

DROP TABLE IF EXISTS users;