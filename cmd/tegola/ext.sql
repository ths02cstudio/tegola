CREATE TABLE
    public.v_areas (
        "version" VARCHAR NOT NULL PRIMARY KEY,
        layer_name VARCHAR NOT NULL,
        start_time TIMESTAMPTZ NOT NULL,
        end_time TIMESTAMPTZ NOT NULL,
        deleted_at TIMESTAMPTZ NULL
    );

CREATE TABLE
    public.c_areas (
        geom_id BIGSERIAL NOT NULL PRIMARY KEY,
        ogc_fid BIGINT NOT NULL,
        ext_properties VARCHAR NULL,
        ext_tags public.hstore NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NULL,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NULL,
        deleted_at TIMESTAMPTZ NULL
    );

-- Select SQL for Test
SELECT
    ma.ogc_fid,
    ma.wkb_geometry,
    ma."geometric_type",
    ma.area_code,
    ma.properties,
    ma.other_tags,
    cat.ext_properties,
    cat.ext_tags,
    cat.created_at
FROM
    (
        SELECT
            *
        FROM
            (
                SELECT
                    *,
                    ROW_NUMBER() OVER (
                        PARTITION BY
                            ca.ogc_fid
                        ORDER BY
                            ca.created_at DESC
                    ) AS rn
                FROM
                    c_areas ca,
                    (
                        SELECT
                            start_time,
                            end_time
                        FROM
                            public.v_areas
                        WHERE
                            "version" = 'v3'
                            AND layer_name = 'areas_polygon'
                    ) AS va
                WHERE
                    ca.created_at >= va.start_time
                    AND ca.created_at <= va.end_time
            ) AS cas
        WHERE
            cas.rn = 1
    ) AS cat
    INNER JOIN public.m_areas ma ON ma.ogc_fid = cat.ogc_fid;