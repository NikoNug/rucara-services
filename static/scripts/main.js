document.addEventListener("DOMContentLoaded", () => {
    const loginForm = document.getElementById('loginForm');
    const signupForm = document.getElementById('signupForm');
    const postsList = document.getElementById('postsList');

    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const email = e.target.email.value;
        const password = e.target.password.value;
        
        const response = await fetch('/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });

        const result = await response.json();
        if (response.ok) {
            alert('Login successful');
            // Store the token, redirect, etc.
        } else {
            alert(result.error);
        }
    });

    signupForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const name = e.target.name.value;
        const email = e.target.email.value;
        const password = e.target.password.value;

        const response = await fetch('/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, email, password }),
        });

        const result = await response.json();
        if (response.ok) {
            alert('Signup successful');
            // Redirect or clear form, etc.
        } else {
            alert(result.error);
        }
    });

    const fetchPosts = async () => {
        console.log('Fetching posts...');
        
        try {
            const response = await fetch('/posts');
            console.log('Response status:', response.status);
            if (!response.ok) {
                throw new Error(`An error occurred: ${response.statusText}`);
            }
            const posts = await response.json();
            console.log('Posts fetched:', posts);

            if (posts.length === 0) {
                postsList.innerHTML = '<p>No posts available.</p>';
                return;
            }

            postsList.innerHTML = '';  // Clear any existing content
            posts.forEach(post => {
                const postItem = document.createElement('div');
                postItem.className = 'post';
                postItem.innerHTML = `
                    <h3>${post.title}</h3>
                    <p>${post.content}</p>
                `;
                postsList.appendChild(postItem);
            });
        } catch (error) {
            console.error('Error:', error);
            postsList.innerHTML = `<p>An error occurred: ${error.message}</p>`;
        }
    };

    fetchPosts();
});
