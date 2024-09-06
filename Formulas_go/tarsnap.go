package main

// Tarsnap - Online backups for the truly paranoid
// Homepage: https://www.tarsnap.com/

import (
	"fmt"
	
	"os/exec"
)

func installTarsnap() {
	// Método 1: Descargar y extraer .tar.gz
	tarsnap_tar_url := "https://www.tarsnap.com/download/tarsnap-autoconf-1.0.40.tgz"
	tarsnap_cmd_tar := exec.Command("curl", "-L", tarsnap_tar_url, "-o", "package.tar.gz")
	err := tarsnap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tarsnap_zip_url := "https://www.tarsnap.com/download/tarsnap-autoconf-1.0.40.tgz"
	tarsnap_cmd_zip := exec.Command("curl", "-L", tarsnap_zip_url, "-o", "package.zip")
	err = tarsnap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tarsnap_bin_url := "https://www.tarsnap.com/download/tarsnap-autoconf-1.0.40.tgz"
	tarsnap_cmd_bin := exec.Command("curl", "-L", tarsnap_bin_url, "-o", "binary.bin")
	err = tarsnap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tarsnap_src_url := "https://www.tarsnap.com/download/tarsnap-autoconf-1.0.40.tgz"
	tarsnap_cmd_src := exec.Command("curl", "-L", tarsnap_src_url, "-o", "source.tar.gz")
	err = tarsnap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tarsnap_cmd_direct := exec.Command("./binary")
	err = tarsnap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: e2fsprogs")
exec.Command("latte", "install", "e2fsprogs")
}
