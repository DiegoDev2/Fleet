package main

// X3270 - IBM 3270 terminal emulator for the X Window System and Windows
// Homepage: http://x3270.bgp.nu/

import (
	"fmt"
	
	"os/exec"
)

func installX3270() {
	// Método 1: Descargar y extraer .tar.gz
	x3270_tar_url := "http://x3270.bgp.nu/download/04.03/suite3270-4.3ga9-src.tgz"
	x3270_cmd_tar := exec.Command("curl", "-L", x3270_tar_url, "-o", "package.tar.gz")
	err := x3270_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	x3270_zip_url := "http://x3270.bgp.nu/download/04.03/suite3270-4.3ga9-src.tgz"
	x3270_cmd_zip := exec.Command("curl", "-L", x3270_zip_url, "-o", "package.zip")
	err = x3270_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	x3270_bin_url := "http://x3270.bgp.nu/download/04.03/suite3270-4.3ga9-src.tgz"
	x3270_cmd_bin := exec.Command("curl", "-L", x3270_bin_url, "-o", "binary.bin")
	err = x3270_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	x3270_src_url := "http://x3270.bgp.nu/download/04.03/suite3270-4.3ga9-src.tgz"
	x3270_cmd_src := exec.Command("curl", "-L", x3270_src_url, "-o", "source.tar.gz")
	err = x3270_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	x3270_cmd_direct := exec.Command("./binary")
	err = x3270_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
