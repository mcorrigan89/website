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

CREATE TABLE IF NOT EXISTS website_content (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,
    locale citext NOT NULL,

    website_display_name TEXT,
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

CREATE TABLE IF NOT EXISTS website_page_content (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_page_id UUID NOT NULL REFERENCES website_page(id) ON DELETE CASCADE,
    locale citext NOT NULL,
    title TEXT,
    subtitle TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS website_component (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_id UUID NOT NULL REFERENCES website(id) ON DELETE CASCADE,
    website_page_id UUID NOT NULL REFERENCES website_page(id) ON DELETE CASCADE,
    sort_key TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS simple_text_component (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    website_component_id UUID NOT NULL REFERENCES website_component(id) ON DELETE CASCADE,
    locale citext NOT NULL,

    content TEXT NOT NULL,

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

ALTER TABLE website_page_content ADD CONSTRAINT page_local_constraint UNIQUE (website_page_id, locale);