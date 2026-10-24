// npm install ws
import WebSocket from 'ws';

const ws = new WebSocket('wss://api.x.ai/v1/realtime?agent_id=agent_LhOy7jXCV7vuxLMq', {
  headers: { Authorization: `Bearer ${process.env.XAI_API_KEY}` },
});

ws.on('open', () => {
  ws.send(JSON.stringify({
    type: 'conversation.item.create',
    item: { type: 'message', role: 'user',
      content: [{ type: 'input_text', text: 'Hello!' }] },
  }));
  ws.send(JSON.stringify({ type: 'response.create' }));
});

ws.on('message', raw => {
  const event = JSON.parse(raw.toString());
  if (event.type === 'response.output_audio_transcript.delta') {
    process.stdout.write(event.delta);
  } else if (event.type === 'response.output_audio.delta') {
    const pcm = Buffer.from(event.delta, 'base64'); // decode and play
  }
});
