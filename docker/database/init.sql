DROP TABLE IF EXISTS credit_line;
CREATE TABLE IF NOT EXISTS credit_line (
    id SERIAL NOT NULL,
    cash_balance decimal(12,2) DEFAULT NULL,
    fail_times bigint DEFAULT NULL,
    founding_types varchar(10) DEFAULT NULL,
    monthly_revenue decimal(12,2) DEFAULT NULL,
    requested_credit_line decimal(12,2) DEFAULT NULL,
    requested_date timestamp DEFAULT NULL,
    valid bool DEFAULT false,
    PRIMARY KEY (id),
    UNIQUE (requested_credit_line)
);

INSERT INTO credit_line (id, cash_balance, fail_times, founding_types, monthly_revenue, requested_credit_line, requested_date, valid) VALUES
    (29, '435.30', 1, 'SME', '10.45', '100.00', '2021-11-30 16:52:04.099000', false);
COMMIT;