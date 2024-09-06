package main

// RedisLeveldb - Redis-protocol compatible frontend to leveldb
// Homepage: https://github.com/KDr2/redis-leveldb

import (
	"fmt"
	
	"os/exec"
)

func installRedisLeveldb() {
	// Método 1: Descargar y extraer .tar.gz
	redisleveldb_tar_url := "https://github.com/KDr2/redis-leveldb/archive/refs/tags/v1.4.tar.gz"
	redisleveldb_cmd_tar := exec.Command("curl", "-L", redisleveldb_tar_url, "-o", "package.tar.gz")
	err := redisleveldb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redisleveldb_zip_url := "https://github.com/KDr2/redis-leveldb/archive/refs/tags/v1.4.zip"
	redisleveldb_cmd_zip := exec.Command("curl", "-L", redisleveldb_zip_url, "-o", "package.zip")
	err = redisleveldb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redisleveldb_bin_url := "https://github.com/KDr2/redis-leveldb/archive/refs/tags/v1.4.bin"
	redisleveldb_cmd_bin := exec.Command("curl", "-L", redisleveldb_bin_url, "-o", "binary.bin")
	err = redisleveldb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redisleveldb_src_url := "https://github.com/KDr2/redis-leveldb/archive/refs/tags/v1.4.src.tar.gz"
	redisleveldb_cmd_src := exec.Command("curl", "-L", redisleveldb_src_url, "-o", "source.tar.gz")
	err = redisleveldb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redisleveldb_cmd_direct := exec.Command("./binary")
	err = redisleveldb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: leveldb")
	exec.Command("latte", "install", "leveldb").Run()
	fmt.Println("Instalando dependencia: libev")
	exec.Command("latte", "install", "libev").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
}
