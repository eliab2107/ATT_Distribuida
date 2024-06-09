import subprocess
from threading import Thread, Lock

# Lista para armazenar os tempos de execução
tempo_por_teste = []
# Lock para sincronizar o acesso à lista compartilhada
lock = Lock()

def client(command):
    result = subprocess.run(command, shell=True, capture_output=True, text=True)
    time_taken = float(result.stdout[:-3])
    with lock:
        tempo_por_teste.append(time_taken)

def run_command_in_threads(command, n):
    threads = []
    
    for i in range(n):
        thread = Thread(target=client, args=(command,))
        threads.append(thread)
        thread.start()
    
    for thread in threads:
        thread.join()

    print("Tempos de execução por teste:", tempo_por_teste)

if __name__ == "__main__":
    command = "go run Sockets\Client\Client.go"
    n = 20
    
    run_command_in_threads(command, n)