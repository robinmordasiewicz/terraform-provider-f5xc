import { createMCPClient } from './dist/utils/mcp-client.js';
import { ResponseFormat } from './dist/types.js';

const client = createMCPClient(ResponseFormat.JSON);

async function checkExamples() {
  const resources = ['tcp_loadbalancer', 'bot_defense_policy', 'certificate'];

  for (const resource of resources) {
    console.log(`\n=== ${resource} ===`);
    try {
      const example = await client.getExample(resource, 'basic');
      if (!example || example.length < 50) {
        console.log('❌ Example is empty or too short');
        console.log('Example:', example);
      } else {
        console.log('✓ Example exists, length:', example.length);
        console.log('First 300 chars:', example.substring(0, 300));
      }
    } catch (e) {
      console.log('❌ ERROR:', e.message);
    }
  }
}

checkExamples();
