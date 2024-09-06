package main

// Socat - SOcket CAT: netcat on steroids
// Homepage: http://www.dest-unreach.org/socat/

import (
	"fmt"
	
	"os/exec"
)

func installSocat() {
	// Método 1: Descargar y extraer .tar.gz
	socat_tar_url := "http://www.dest-unreach.org/socat/download/socat-1.8.0.1.tar.gz"
	socat_cmd_tar := exec.Command("curl", "-L", socat_tar_url, "-o", "package.tar.gz")
	err := socat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	socat_zip_url := "http://www.dest-unreach.org/socat/download/socat-1.8.0.1.zip"
	socat_cmd_zip := exec.Command("curl", "-L", socat_zip_url, "-o", "package.zip")
	err = socat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	socat_bin_url := "http://www.dest-unreach.org/socat/download/socat-1.8.0.1.bin"
	socat_cmd_bin := exec.Command("curl", "-L", socat_bin_url, "-o", "binary.bin")
	err = socat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	socat_src_url := "http://www.dest-unreach.org/socat/download/socat-1.8.0.1.src.tar.gz"
	socat_cmd_src := exec.Command("curl", "-L", socat_src_url, "-o", "source.tar.gz")
	err = socat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	socat_cmd_direct := exec.Command("./binary")
	err = socat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
