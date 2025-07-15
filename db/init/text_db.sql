-- Tabla de cursos
CREATE TABLE courses (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description VARCHAR(300) NOT NULL,
    image_url VARCHAR(500),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Tabla de secciones
CREATE TABLE sections (
    id BIGSERIAL PRIMARY KEY,
    course_id BIGINT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    "order" INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Tabla de artículos
CREATE TABLE articles (
    id BIGSERIAL PRIMARY KEY,
    section_id BIGINT NOT NULL REFERENCES sections(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Índices para optimización de consultas
CREATE INDEX idx_sections_course_id ON sections(course_id);
CREATE INDEX idx_articles_section_id ON articles(section_id);
CREATE INDEX idx_courses_title ON courses(title);
CREATE INDEX idx_sections_title ON sections(title);
CREATE INDEX idx_articles_title ON articles(title);


DO $$
DECLARE
  course_id BIGINT;
  section1_id BIGINT;
  section2_id BIGINT;
  section3_id BIGINT;
BEGIN
  -- Curso
  INSERT INTO courses (title, description, image_url)
  VALUES ('Basic Foundations', 'Learn the basics of programming from scratch.', '/replace')
  RETURNING id INTO course_id;

  -- Sección 1
  INSERT INTO sections (course_id, title, description, "order")
  VALUES (course_id, 'Programming Logic', 'Understand what programming is and how languages work.', 1)
  RETURNING id INTO section1_id;

  INSERT INTO articles (section_id, title, content) VALUES
  (section1_id, 'What is Programming?',
  '<p>Programming logic is the foundation of algorithmic thinking. It enables the creation of step-by-step solutions to solve computational problems.</p>
   <p>The goal is to develop analysis skills, process structuring, and problem-solving using logical sequences.</p>
   <p>Example in pseudocode to check if a number is even or odd:</p>
   <pre><code>Start
  Read number
  If number MOD 2 = 0 then
    Print "The number is even"
  Else
    Print "The number is odd"
End</code></pre>
   <p>These types of exercises help you understand conditionals and loops before learning a specific programming language.</p>'),

  (section1_id, 'Programming Languages',
  '<h2>Popular Languages</h2><ul><li>Python</li><li>JavaScript</li><li>Java</li><li>C++</li></ul><p>Each language has its own strengths depending on the type of project.</p>');

  -- Sección 2
  INSERT INTO sections (course_id, title, description, "order")
  VALUES (course_id, 'Algorithms and Computer Fundamentals', 'Core concepts to understand how computers process data and instructions.', 2)
  RETURNING id INTO section2_id;

  INSERT INTO articles (section_id, title, content) VALUES
  (section2_id, 'Algorithms and Data Structures',
  '<p>An algorithm is a finite series of steps to solve a problem. Data structures are ways to organize information efficiently for processing.</p>
   <p>Common data structures include arrays, linked lists, stacks, queues, trees, and graphs.</p>
   <p>Example: summing the elements of an array in JavaScript:</p>
   <pre><code>function sumArray(arr) {
  let sum = 0;
  for (let i = 0; i &lt; arr.length; i++) {
    sum += arr[i];
  }
  return sum;
}</code></pre>
   <p>This shows how to use iteration to manipulate collections of data.</p>'),

  (section2_id, 'Computer Fundamentals',
  '<p>This topic introduces the essential components of a computer and how they work together. It includes hardware, software, operating systems, binary data, and basic architecture.</p>
   <p>It''s important to understand how a computer interprets instructions and represents data internally.</p>
   <p>Example of decimal to binary conversion:</p>
   <pre><code>Decimal: 25
Binary: 11001</code></pre>
   <p>You''ll also explore how programs run, the function of the CPU, memory, and input/output devices.</p>');

  -- Sección 3
  INSERT INTO sections (course_id, title, description, "order")
  VALUES (course_id, 'Applied Programming Concepts', 'Learn object-oriented and web programming principles.', 3)
  RETURNING id INTO section3_id;

  INSERT INTO articles (section_id, title, content) VALUES
  (section3_id, 'Object-Oriented Programming',
  '<p>Object-Oriented Programming (OOP) is a paradigm that structures code around "objects" which contain data (attributes) and behaviors (methods).</p>
   <p>Key principles include classes, inheritance, encapsulation, and polymorphism.</p>
   <p>Example in JavaScript:</p>
   <pre><code>class Animal {
  constructor(name) {
    this.name = name;
  }

  speak() {
    console.log(`${this.name} makes a sound.`);
  }
}

const dog = new Animal("Buddy");
dog.speak();</code></pre>
   <p>OOP helps build modular, reusable, and maintainable code.</p>'),

  (section3_id, 'Introduction to Web Development',
  '<p>Web development involves building websites and applications that run in a browser. It usually includes frontend (interface) and backend (server logic).</p>
   <p>Core technologies are: HTML (structure), CSS (style), and JavaScript (behavior).</p>
   <p>Basic HTML document example:</p>
   <pre><code>&lt;!DOCTYPE html&gt;
&lt;html&gt;
  &lt;head&gt;
    &lt;title&gt;My First Website&lt;/title&gt;
  &lt;/head&gt;
  &lt;body&gt;
    &lt;h1&gt;Hello World!&lt;/h1&gt;
    &lt;p&gt;This is my first website using HTML.&lt;/p&gt;
  &lt;/body&gt;
&lt;/html&gt;</code></pre>
   <p>As you progress, you''ll learn how to add styles, interactivity, and backend integration to build full-featured web applications.</p>');

END;
$$;
