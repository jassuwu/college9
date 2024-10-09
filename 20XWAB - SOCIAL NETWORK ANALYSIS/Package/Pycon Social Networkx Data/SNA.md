# **Influence Maximization in Social Networks Using Game of Thrones Network**

---

## **Introduction**

- We are analyzing the **Game of Thrones** character network to simulate **information spread** and **influence maximization**.
- The network is derived from the co-occurrences of characters in the books.
- We will explore how information spreads, identify key influencers, and simulate different strategies to maximize influence.

---

## **Dataset Overview**

- The dataset represents character interactions from **A Storm of Swords** (GoT series).
- **Nodes**: Characters.
- **Edges**: Connections between characters, weighted by how frequently they interact.

---

## **Key Concepts and Terms**

### 1. **Seeding Strategies**
- Refers to selecting **initial nodes** (seeds) in a network to begin spreading influence.
- Example: Choosing **highly connected characters** like Jon Snow to start spreading information.

---

### 2. **Centrality Measures**
These are used to rank nodes (characters) based on their importance:

- **Degree Centrality**: Measures the number of direct connections a character has.
  - Example: A character with many relationships has high degree centrality.
  
- **PageRank**: Ranks characters based on both quantity and quality of their connections (like Google’s web ranking).
  - Example: Characters connected to influential people have high PageRank.
  
- **Betweenness Centrality**: Measures how often a character lies on the shortest path between other characters.
  - Example: Characters like **Varys** who connect different groups have high betweenness.

---

### 3. **Subgraph**
- A smaller part of the network that focuses on key characters with more interactions.
- Example: The subgraph in this study focuses on characters with **weighted degrees** above 200.

---

### 4. **Weighted Degree**
- Number of connections a character has, weighted by the frequency of interaction.
- Example: Characters who interact more frequently have higher weighted degrees.

---

### 5. **Independent Cascade Model (IC)**
- A simulation model to see how **influence** or **information** spreads through the network.
- Each infected node has a chance to infect its neighbors, based on edge weights.
- Example: If **Jon Snow** starts as the seed, he might influence neighboring characters in subsequent time steps.

---

### 6. **Edge Weight**
- Represents the strength or frequency of the interaction between two characters.
- Example: **Tyrion** and **Jaime Lannister** may have a high edge weight because of their frequent interactions.

---

### 7. **Spring Layout**
- A method for arranging nodes in a graph visually.
- Treats edges like springs: connected nodes are drawn closer together.
- Example: Characters that interact more are positioned closer to each other in the visualization.

---

### 8. **Information Spread**
- Refers to how information or influence propagates from one character to another in the network.
- Example: Starting from **Jon Snow**, the information spreads across the GoT network over time.

---

### 9. **Infection Times**
- Time at which a node (character) becomes infected or influenced in the simulation.
- Example: **Arya Stark** might be infected at time `t=3`, meaning she was influenced at the third step of the simulation.

---

### 10. **Force-Directed Algorithm**
- Used to arrange the nodes and edges of a graph.
- Nodes repel each other, while edges pull connected nodes closer.
- Example: Used in the **spring layout** for positioning characters based on their interactions.

---

### 11. **Influence Maximization**
- The process of selecting the best nodes (characters) to start spreading influence in order to maximize the number of influenced nodes.
- Example: Using **centrality measures** to find the most influential characters.

---

### 12. **Brute Force Search**
- A method where all possible combinations are tried to find the best solution.
- Example: In the GoT network, brute force search is used to find the **best pair of characters** to start the infection.

---

### 13. **Aggregating Edges**
- Combining multiple edges between the same nodes and summing their weights.
- Example: If two characters appear together multiple times, their edge is aggregated.

---

### 14. **Histogram**
- A plot that shows the frequency distribution of a particular value.
- Example: The histogram shows the distribution of **weighted degrees** among characters in the GoT network.

---

## **Steps in the Code**

### 1. **Loading the Data**
- The dataset is loaded and combined from 5 books.
- Edges between characters are summed based on the number of interactions.

---

### 2. **Creating the Graph**
- A **network graph** is created using the edge list, connecting characters based on interaction frequency.

---

### 3. **Weighted Degree Distribution**
- Calculate the **weighted degree** of each character and plot the distribution.

---

### 4. **Creating the Subgraph**
- Focus on the most influential characters by creating a subgraph of those with high weighted degrees.

---

### 5. **Simulating Information Spread**
- Using the **Independent Cascade Model**, we simulate how information spreads across the network starting from certain characters.
- Visualize the network at each step to track the spread.

---

### 6. **Seeding Strategies and Influence Maximization**
- Test different **seeding strategies** by selecting influential characters based on centrality measures.
- Simulate the spread for different **seeding budgets** and compare results.

---

### 7. **Finding the Best Pair**
- Use **brute force** to find the best **pair of characters** to start the spread when the budget is limited to 2.

---

## **Conclusion**

- **Goal**: Identify the most influential characters in the GoT network and optimize information spread.
- **Key Techniques**: Centrality measures, Independent Cascade Model, and seeding strategies.
- **Result**: Using centrality measures and simulations, we can identify optimal characters for influence maximization in the network.
