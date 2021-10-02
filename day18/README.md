--- Day 18: Operation Order ---
-------------------------------

As you look out the window and notice a heavily-forested continent slowly appear over the horizon, you are interrupted by the child sitting next to you. They're curious if you could help them with their <span title="Or "maths", if you have more than one.">math</span> homework.

Unfortunately, it seems like this "math" [follows different rules](https://www.youtube.com/watch?v=3QtRK7Y2pPU&t=15) than you remember.

The homework (your puzzle input) consists of a series of expressions that consist of addition (`+`), multiplication (`*`), and parentheses (`(...)`). Just like normal math, parentheses indicate that the expression inside must be evaluated before it can be used by the surrounding expression. Addition still finds the sum of the numbers on both sides of the operator, and multiplication still finds the product.

However, the rules of *operator precedence* have changed. Rather than evaluating multiplication before addition, the operators have the *same precedence*, and are evaluated left-to-right regardless of the order in which they appear.

For example, the steps to evaluate the expression `1 + 2 * 3 + 4 * 5 + 6` are as follows:

```
<em>1 + 2</em> * 3 + 4 * 5 + 6
  <em>3   * 3</em> + 4 * 5 + 6
      <em>9   + 4</em> * 5 + 6
         <em>13   * 5</em> + 6
             <em>65   + 6</em>
                 <em>71</em>

```

Parentheses can override this order; for example, here is what happens if parentheses are added to form `1 + (2 * 3) + (4 * (5 + 6))`:

```
1 + <em>(2 * 3)</em> + (4 * (5 + 6))
<em>1 +    6</em>    + (4 * (5 + 6))
     7      + (4 * <em>(5 + 6)</em>)
     7      + <em>(4 *   11   )</em>
     <em>7      +     44</em>
            <em>51</em>

```

Here are a few more examples:

- `2 * 3 + (4 * 5)` becomes *`26`*.
- `5 + (8 * 3 + 9 + 3 * 4 * 3)` becomes *`437`*.
- `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))` becomes *`12240`*.
- `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2` becomes *`13632`*.

Before you can help with the homework, you need to understand it yourself. *Evaluate the expression on each line of the homework; what is the sum of the resulting values?*

-----

--- Part Two ---
----------------

You manage to answer the child's questions and they finish part 1 of their homework, but get stuck when they reach the next section: *advanced* math.

Now, addition and multiplication have *different* precedence levels, but they're not the ones you're familiar with. Instead, addition is evaluated *before* multiplication.

For example, the steps to evaluate the expression `1 + 2 * 3 + 4 * 5 + 6` are now as follows:

```
<em>1 + 2</em> * 3 + 4 * 5 + 6
  3   * <em>3 + 4</em> * 5 + 6
  3   *   7   * <em>5 + 6</em>
  <em>3   *   7</em>   *  11
     <em>21       *  11</em>
         <em>231</em>

```

Here are the other examples from above:

- `1 + (2 * 3) + (4 * (5 + 6))` still becomes *`51`*.
- `2 * 3 + (4 * 5)` becomes *`46`*.
- `5 + (8 * 3 + 9 + 3 * 4 * 3)` becomes *`1445`*.
- `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))` becomes *`669060`*.
- `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2` becomes *`23340`*.

*What do you get if you add up the results of evaluating the homework problems using these new rules?*