<?php

namespace Tests\Feature;

use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Foundation\Testing\WithFaker;
use Tests\TestCase;

class LoginTest extends TestCase
{
    public function test_login_invalid_params(): void
    {
        $response = $this->post('/api/login');
        $response->assertStatus(400);
    }
}
