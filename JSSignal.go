package gsc

func JSSignal() string {
	return Script(``, `
		function jsSignal(initialValue) {
			let value = initialValue;
			let subscribers = [];

			function get() {
				return value;
			}

			function set(newValue) {
				if (value !== newValue) {
				value = newValue;
				subscribers.forEach(callback => callback(value));
				}
			}

			function subscribe(callback) {
				subscribers.push(callback);
				return () => {
				subscribers = subscribers.filter(sub => sub !== callback);
				};
			}

			return { get, set, subscribe };

		}
	`)
}
