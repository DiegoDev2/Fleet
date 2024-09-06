package main

// Tevent - Event system based on the talloc memory management library
// Homepage: https://tevent.samba.org

import (
	"fmt"
	
	"os/exec"
)

func installTevent() {
	// Método 1: Descargar y extraer .tar.gz
	tevent_tar_url := "https://www.samba.org/ftp/tevent/tevent-0.16.1.tar.gz"
	tevent_cmd_tar := exec.Command("curl", "-L", tevent_tar_url, "-o", "package.tar.gz")
	err := tevent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tevent_zip_url := "https://www.samba.org/ftp/tevent/tevent-0.16.1.zip"
	tevent_cmd_zip := exec.Command("curl", "-L", tevent_zip_url, "-o", "package.zip")
	err = tevent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tevent_bin_url := "https://www.samba.org/ftp/tevent/tevent-0.16.1.bin"
	tevent_cmd_bin := exec.Command("curl", "-L", tevent_bin_url, "-o", "binary.bin")
	err = tevent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tevent_src_url := "https://www.samba.org/ftp/tevent/tevent-0.16.1.src.tar.gz"
	tevent_cmd_src := exec.Command("curl", "-L", tevent_src_url, "-o", "source.tar.gz")
	err = tevent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tevent_cmd_direct := exec.Command("./binary")
	err = tevent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmocka")
	exec.Command("latte", "install", "cmocka").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: talloc")
	exec.Command("latte", "install", "talloc").Run()
}
