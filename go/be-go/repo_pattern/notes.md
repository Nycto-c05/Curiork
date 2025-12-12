Alright, let’s break this down **super simply** and **super concretely**, because you’re basically asking two things:

1. **“Why can’t I just do this?”**

   ```go
   type Storage struct {
       Posts interface { Get(); Post(); Put() }
       Comments interface { Create(); Post() }
   }
   ```

2. **“If I *did* do it, what would Get() / Post() actually look like?”**

Let’s answer both.
No BS. No overengineering. Just the truth.

---

# ✅ **1. Why this is a bad idea (but legal Go)**

Your struct:

```go
type Storage struct {
    Posts interface {
        Get()
        Post()
        Put()
    }

    Comments interface {
        Create()
        Post()
    }
}
```

**is syntactically valid**, but it creates 2 brand-new, anonymous interface *types*.

So:
`Posts` has **its own private anonymous interface type**
`Comments` has **its own private anonymous interface type**

Even if they had identical methods, Go would say:

> These two are completely different types.

This means:

* You cannot reuse these types
* You cannot refer to them
* You cannot mock them
* You cannot pass them to functions
* You cannot check if something implements them
* You cannot wrap them (logging/caching)
* You cannot extend them

Basically: **You trap your interfaces inside the struct, where they can’t be used by anything else.**

This is why Go developers avoid this pattern except in tiny one-off cases.

---

# ✅ **2. HOW WOULD YOU IMPLEMENT IT IF YOU REALLY WANTED TO?**

Okay — here’s EXACTLY what you want to see:

**Concrete structs that implement those anonymous interfaces.**

Say you do:

```go
type Storage struct {
	Posts interface {
		Get()
		Post()
		Put()
	}

	Comments interface {
		Create()
		Post()
	}
}
```

Now you need something like:

---

## 🟦 Posts Implementation

```go
type DBPosts struct{}

func (DBPosts) Get() {
    fmt.Println("Posts.Get() called")
}

func (DBPosts) Post() {
    fmt.Println("Posts.Post() called")
}

func (DBPosts) Put() {
    fmt.Println("Posts.Put() called")
}
```

---

## 🟨 Comments Implementation

```go
type DBComments struct{}

func (DBComments) Create() {
    fmt.Println("Comments.Create() called")
}

func (DBComments) Post() {
    fmt.Println("Comments.Post() called")
}
```

---

## 🟩 Now you construct your Storage:

```go
store := Storage{
    Posts:    DBPosts{},
    Comments: DBComments{},
}
```

YES — this compiles.

---

## 🟧 And using it:

```go
store.Posts.Get()
store.Posts.Post()
store.Posts.Put()

store.Comments.Create()
store.Comments.Post()
```

ALL OF THESE WORK.

---

# 🔥 SO WHAT’S THE PROBLEM THEN?

The problem is NOT that this doesn’t compile.

The problem is that:

### ❌ You can’t write functions that accept these interfaces

Like:

```go
func HandlePosts(p ???) {}
```

There is NO NAME for the type of `Posts`.

You’d have to rewrite the ENTIRE interface every single time:

```go
func HandlePosts(p interface { Get(); Post(); Put() }) {}
```

If you change ONE method, you must update it EVERYWHERE.

This becomes a clusterfuck fast.

---

# 🔥 TL;DR

### ✔️ Your anonymous-interface Storage struct works

### ❌ But it becomes **unusable** outside the struct

### ❌ And you can’t reference the interface type anywhere

### ✔️ You CAN implement it like I showed above

### ❌ But it’s bad architecture because Go treats each anonymous interface as a separate unique type

---

# If you want the **correct, idiomatic**, reusable version:

I can rewrite your exact struct with named interfaces **so everything makes sense**.

Just say *“show me the proper version.”*
