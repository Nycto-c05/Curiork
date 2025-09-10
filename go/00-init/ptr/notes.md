Great instinct — this is one of the subtle things in Go that trips people up.

---

### 🔑 In Go:

* Methods can have **value receivers** (`func (r Rectangle)`) or **pointer receivers** (`func (r *Rectangle)`).
* **They are not “always by reference.”** It depends on how you declare the receivey was scaled, original unchanged
}
```
                                                                            `go
type Rectangle struct {
    Width, Height floa                                                                           t64
}

func (r *Rectangle) Scale(factor float64) {unc main() {
    rect := Rectangle{2, 3}
    rect.Scale(2)       // Go automatically takes &rect
  ---

### ⚡ Why the confusion?

When you call `rect.Scale(2)` on a pointer receiver, Go **automatically takes the address** (`&rect`) behind the scenes. That’s why it *feels* like structs are always byerence**. You decide with `T` (value) vs `*T` (poD for when to use value vs pointer receivers?
o you want me to break down the **rule of thumb**ilues convenient by auto-taking the address.

---

nter).
Go just makes calling pointer methods on va So:
Struct methods are **not automatically by refreceiver**, Go passes the actual address.

---

👉, Go passes a copy.
* If defined with a **pointer 

* If method is defined with a **value receiver** reference — but really it’s just syntactic sugar.```

Now `r` points to the original `Rectangle`.

  fmt.Println(rect)   // {4 6}, original changed
}
                                                                           
    r.Width *= factor
    r.Height *= factor
}

f
Here `r` is a **copy**. The original `rect` stays the same.

---

### 📌 Example with **pointer receiver**

``r.

---

### 📌 Example with **value receiver**

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{2, 3}
    rect.Scale(2)
    fmt.Println(rect) // still {2 3}, cop
