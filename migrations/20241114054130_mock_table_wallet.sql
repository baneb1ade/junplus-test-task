-- +goose Up
-- +goose StatementBegin
INSERT INTO wallet (id, balance) VALUES ('a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29', 150.75);
INSERT INTO wallet (id, balance) VALUES ('bbd9c3f1-8a5f-4f3e-87e6-9c8b4a9d69c0', 2000.00);
INSERT INTO wallet (id, balance) VALUES ('cc7a9d85-f728-4c44-b55b-34e354f5937a', 500.50);
INSERT INTO wallet (id, balance) VALUES ('dde3f8e2-91a7-47fc-b09e-4f52934912a8', 750.25);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM wallet WHERE id = 'a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29';
DELETE FROM wallet WHERE id = 'bbd9c3f1-8a5f-4f3e-87e6-9c8b4a9d69c0';
DELETE FROM wallet WHERE id = 'cc7a9d85-f728-4c44-b55b-34e354f5937a';
DELETE FROM wallet WHERE id = 'dde3f8e2-91a7-47fc-b09e-4f52934912a8';
-- +goose StatementEnd
