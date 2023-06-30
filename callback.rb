def process_callbacks(callbacks)
    callbacks.each(&:call)
  end
  
  # Example callback lambdas
  callback1 = lambda { puts 'Callback 1' }
  callback2 = lambda { puts 'Callback 2' }
  
  # Create an array of 100,000 callbacks
  callbacks = [callback1] * 50000 + [callback2] * 50000
  
  # Process the callbacks
  process_callbacks(callbacks)