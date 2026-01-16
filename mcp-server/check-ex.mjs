import { createMCPClient } from './tests/utils/mcp-client.js';
import { ResponseFormat } from './src/types.js';

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
        const firstLine = example.split('\n')[0];
        console.log('First line:', firstLine);
        // Check if it has resource blocks
        const hasResource = example.includes('resource "f5xc_');
        console.log('Has resource blocks:', hasResource);
      }
    } catch (e) {
      console.log('❌ ERROR:', e.message);
    }
  }
}

checkExamples();
