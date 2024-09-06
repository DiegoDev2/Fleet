package main

// SpawnFcgi - Spawn FastCGI processes
// Homepage: https://redmine.lighttpd.net/projects/spawn-fcgi

import (
	"fmt"
	
	"os/exec"
)

func installSpawnFcgi() {
	// Método 1: Descargar y extraer .tar.gz
	spawnfcgi_tar_url := "https://www.lighttpd.net/download/spawn-fcgi-1.6.5.tar.gz"
	spawnfcgi_cmd_tar := exec.Command("curl", "-L", spawnfcgi_tar_url, "-o", "package.tar.gz")
	err := spawnfcgi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spawnfcgi_zip_url := "https://www.lighttpd.net/download/spawn-fcgi-1.6.5.zip"
	spawnfcgi_cmd_zip := exec.Command("curl", "-L", spawnfcgi_zip_url, "-o", "package.zip")
	err = spawnfcgi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spawnfcgi_bin_url := "https://www.lighttpd.net/download/spawn-fcgi-1.6.5.bin"
	spawnfcgi_cmd_bin := exec.Command("curl", "-L", spawnfcgi_bin_url, "-o", "binary.bin")
	err = spawnfcgi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spawnfcgi_src_url := "https://www.lighttpd.net/download/spawn-fcgi-1.6.5.src.tar.gz"
	spawnfcgi_cmd_src := exec.Command("curl", "-L", spawnfcgi_src_url, "-o", "source.tar.gz")
	err = spawnfcgi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spawnfcgi_cmd_direct := exec.Command("./binary")
	err = spawnfcgi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
