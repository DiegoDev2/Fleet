package main

// RedisAT62 - Persistent key-value database, with built-in net interface
// Homepage: https://redis.io/

import (
	"fmt"
	
	"os/exec"
)

func installRedisAT62() {
	// Método 1: Descargar y extraer .tar.gz
	redisat62_tar_url := "https://download.redis.io/releases/redis-6.2.14.tar.gz"
	redisat62_cmd_tar := exec.Command("curl", "-L", redisat62_tar_url, "-o", "package.tar.gz")
	err := redisat62_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redisat62_zip_url := "https://download.redis.io/releases/redis-6.2.14.zip"
	redisat62_cmd_zip := exec.Command("curl", "-L", redisat62_zip_url, "-o", "package.zip")
	err = redisat62_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redisat62_bin_url := "https://download.redis.io/releases/redis-6.2.14.bin"
	redisat62_cmd_bin := exec.Command("curl", "-L", redisat62_bin_url, "-o", "binary.bin")
	err = redisat62_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redisat62_src_url := "https://download.redis.io/releases/redis-6.2.14.src.tar.gz"
	redisat62_cmd_src := exec.Command("curl", "-L", redisat62_src_url, "-o", "source.tar.gz")
	err = redisat62_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redisat62_cmd_direct := exec.Command("./binary")
	err = redisat62_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
