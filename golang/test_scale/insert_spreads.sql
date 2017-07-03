delete from spreads;

insert into spreads
(created_at, updated_at, deleted_at, fx_provider, sell_currency_id, buy_currency_id, spread1, spread2)
VALUES
(CURRENT_TIMESTAMP, null, null, 'Scale', 'SGD', 'USD', 0.01, 0.015);

insert into spreads
(created_at, updated_at, deleted_at, fx_provider, sell_currency_id, buy_currency_id, spread1, spread2)
VALUES
(CURRENT_TIMESTAMP, null, null, 'Scale', 'USD', 'SGD', 0.01, 0.015);

insert into spreads
(created_at, updated_at, deleted_at, fx_provider, sell_currency_id, buy_currency_id, spread1, spread2)
VALUES
(CURRENT_TIMESTAMP, null, null, 'Scale', 'SGD', 'HKD', 0.005, 0.008);

insert into spreads
(created_at, updated_at, deleted_at, fx_provider, sell_currency_id, buy_currency_id, spread1, spread2)
VALUES
(CURRENT_TIMESTAMP, null, null, 'Scale', 'HKD', 'SGD', 0.005, 0.008);

commit;
