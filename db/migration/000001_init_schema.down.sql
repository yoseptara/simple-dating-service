DROP INDEX IF EXISTS "swipes"."index_swipes_on_user_id_target_id_direction";
DROP INDEX IF EXISTS "swipes"."index_swipes_on_user_id_swipe_date";
ALTER TABLE "swipes" DROP CONSTRAINT IF EXISTS "swipes_user_id_fkey";
DROP TABLE IF EXISTS "swipes" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;
