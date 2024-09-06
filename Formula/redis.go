package main

// Redis - Persistent key-value database, with built-in net interface
// Homepage: https://redis.io/

import (
	"fmt"
	
	"os/exec"
)

func installRedis() {
	// Método 1: Descargar y extraer .tar.gz
	redis_tar_url := "https://download.redis.io/releases/redis-7.2.5.tar.gz"
	redis_cmd_tar := exec.Command("curl", "-L", redis_tar_url, "-o", "package.tar.gz")
	err := redis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redis_zip_url := "https://download.redis.io/releases/redis-7.2.5.zip"
	redis_cmd_zip := exec.Command("curl", "-L", redis_zip_url, "-o", "package.zip")
	err = redis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redis_bin_url := "https://download.redis.io/releases/redis-7.2.5.bin"
	redis_cmd_bin := exec.Command("curl", "-L", redis_bin_url, "-o", "binary.bin")
	err = redis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redis_src_url := "https://download.redis.io/releases/redis-7.2.5.src.tar.gz"
	redis_cmd_src := exec.Command("curl", "-L", redis_src_url, "-o", "source.tar.gz")
	err = redis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redis_cmd_direct := exec.Command("./binary")
	err = redis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
