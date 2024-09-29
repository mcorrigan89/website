CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS website (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  
  handle TEXT NOT NULL,
  default_locale citext NOT NULL,

  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  version integer NOT NULL DEFAULT 1
);


CREATE TABLE IF NOT EXISTS website_styles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS palette (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_styles_id UUID NOT NULL REFERENCES website_styles(id) ON DELETE CASCADE,

    color_one TEXT NOT NULL,
    color_two TEXT NOT NULL,
    color_three TEXT NOT NULL,
    color_four TEXT NOT NULL,
    color_five TEXT NOT NULL,
    color_six TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_content (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,
    locale citext NOT NULL,

    website_display_name TEXT NOT NULL,
    website_display_description TEXT,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_page (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,
    url_slug citext NOT NULL,
    sort_key TEXT NOT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_config (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,
    
    default_page_id UUID NOT NULL REFERENCES website_page(id) ON DELETE CASCADE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_page_content (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_page_id UUID NOT NULL REFERENCES website_page(id) ON DELETE CASCADE,
    locale citext NOT NULL,
    title TEXT NOT NULL,
    subtitle TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_section (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,
    website_page_id UUID NOT NULL REFERENCES website_page(id) ON DELETE CASCADE,
    sort_key TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_section_display (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_section_id UUID NOT NULL REFERENCES website_section(id) ON DELETE CASCADE,

    row_count INTEGER NOT NULL,
    image_id UUID,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_component (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,
    website_section_id UUID NOT NULL REFERENCES website_section(id) ON DELETE CASCADE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_component_display (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_component_id UUID NOT NULL REFERENCES website_component(id) ON DELETE CASCADE,

    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    x_coordinate INTEGER NOT NULL,
    y_coordinate INTEGER NOT NULL,

    mobile_height INTEGER,
    mobile_width INTEGER,
    mobile_x_coordinate INTEGER,
    mobile_y_coordinate INTEGER,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS text_component (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_component_id UUID NOT NULL REFERENCES website_component(id) ON DELETE CASCADE,
    locale citext NOT NULL,

    content_json JSON,
    content_html TEXT,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS image_component (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_component_id UUID NOT NULL REFERENCES website_component(id) ON DELETE CASCADE,

    image_id UUID NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

ALTER TABLE website_page_content ADD CONSTRAINT page_locale_constraint UNIQUE (website_page_id, locale);
