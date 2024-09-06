package main

// Fwknop - Single Packet Authorization and Port Knocking
// Homepage: https://www.cipherdyne.org/fwknop/

import (
	"fmt"
	
	"os/exec"
)

func installFwknop() {
	// Método 1: Descargar y extraer .tar.gz
	fwknop_tar_url := "https://www.cipherdyne.org/fwknop/download/fwknop-2.6.11.tar.gz"
	fwknop_cmd_tar := exec.Command("curl", "-L", fwknop_tar_url, "-o", "package.tar.gz")
	err := fwknop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fwknop_zip_url := "https://www.cipherdyne.org/fwknop/download/fwknop-2.6.11.zip"
	fwknop_cmd_zip := exec.Command("curl", "-L", fwknop_zip_url, "-o", "package.zip")
	err = fwknop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fwknop_bin_url := "https://www.cipherdyne.org/fwknop/download/fwknop-2.6.11.bin"
	fwknop_cmd_bin := exec.Command("curl", "-L", fwknop_bin_url, "-o", "binary.bin")
	err = fwknop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fwknop_src_url := "https://www.cipherdyne.org/fwknop/download/fwknop-2.6.11.src.tar.gz"
	fwknop_cmd_src := exec.Command("curl", "-L", fwknop_src_url, "-o", "source.tar.gz")
	err = fwknop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fwknop_cmd_direct := exec.Command("./binary")
	err = fwknop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: libassuan")
exec.Command("latte", "install", "libassuan")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: iptables")
exec.Command("latte", "install", "iptables")
}
