package main

// Pass - Password manager
// Homepage: https://www.passwordstore.org/

import (
	"fmt"
	
	"os/exec"
)

func installPass() {
	// Método 1: Descargar y extraer .tar.gz
	pass_tar_url := "https://git.zx2c4.com/password-store/snapshot/password-store-1.7.4.tar.xz"
	pass_cmd_tar := exec.Command("curl", "-L", pass_tar_url, "-o", "package.tar.gz")
	err := pass_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pass_zip_url := "https://git.zx2c4.com/password-store/snapshot/password-store-1.7.4.tar.xz"
	pass_cmd_zip := exec.Command("curl", "-L", pass_zip_url, "-o", "package.zip")
	err = pass_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pass_bin_url := "https://git.zx2c4.com/password-store/snapshot/password-store-1.7.4.tar.xz"
	pass_cmd_bin := exec.Command("curl", "-L", pass_bin_url, "-o", "binary.bin")
	err = pass_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pass_src_url := "https://git.zx2c4.com/password-store/snapshot/password-store-1.7.4.tar.xz"
	pass_cmd_src := exec.Command("curl", "-L", pass_src_url, "-o", "source.tar.gz")
	err = pass_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pass_cmd_direct := exec.Command("./binary")
	err = pass_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-getopt")
exec.Command("latte", "install", "gnu-getopt")
	fmt.Println("Instalando dependencia: gnupg")
exec.Command("latte", "install", "gnupg")
	fmt.Println("Instalando dependencia: qrencode")
exec.Command("latte", "install", "qrencode")
	fmt.Println("Instalando dependencia: tree")
exec.Command("latte", "install", "tree")
}
