package main

// Libprelude - Universal Security Information & Event Management (SIEM) system
// Homepage: https://www.prelude-siem.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibprelude() {
	// Método 1: Descargar y extraer .tar.gz
	libprelude_tar_url := "https://www.prelude-siem.org/attachments/download/1395/libprelude-5.2.0.tar.gz"
	libprelude_cmd_tar := exec.Command("curl", "-L", libprelude_tar_url, "-o", "package.tar.gz")
	err := libprelude_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libprelude_zip_url := "https://www.prelude-siem.org/attachments/download/1395/libprelude-5.2.0.zip"
	libprelude_cmd_zip := exec.Command("curl", "-L", libprelude_zip_url, "-o", "package.zip")
	err = libprelude_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libprelude_bin_url := "https://www.prelude-siem.org/attachments/download/1395/libprelude-5.2.0.bin"
	libprelude_cmd_bin := exec.Command("curl", "-L", libprelude_bin_url, "-o", "binary.bin")
	err = libprelude_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libprelude_src_url := "https://www.prelude-siem.org/attachments/download/1395/libprelude-5.2.0.src.tar.gz"
	libprelude_cmd_src := exec.Command("curl", "-L", libprelude_src_url, "-o", "source.tar.gz")
	err = libprelude_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libprelude_cmd_direct := exec.Command("./binary")
	err = libprelude_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
