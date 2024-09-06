package main

// Heimdal - Free Kerberos 5 implementation
// Homepage: https://www.h5l.org

import (
	"fmt"
	
	"os/exec"
)

func installHeimdal() {
	// Método 1: Descargar y extraer .tar.gz
	heimdal_tar_url := "https://github.com/heimdal/heimdal/releases/download/heimdal-7.8.0/heimdal-7.8.0.tar.gz"
	heimdal_cmd_tar := exec.Command("curl", "-L", heimdal_tar_url, "-o", "package.tar.gz")
	err := heimdal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	heimdal_zip_url := "https://github.com/heimdal/heimdal/releases/download/heimdal-7.8.0/heimdal-7.8.0.zip"
	heimdal_cmd_zip := exec.Command("curl", "-L", heimdal_zip_url, "-o", "package.zip")
	err = heimdal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	heimdal_bin_url := "https://github.com/heimdal/heimdal/releases/download/heimdal-7.8.0/heimdal-7.8.0.bin"
	heimdal_cmd_bin := exec.Command("curl", "-L", heimdal_bin_url, "-o", "binary.bin")
	err = heimdal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	heimdal_src_url := "https://github.com/heimdal/heimdal/releases/download/heimdal-7.8.0/heimdal-7.8.0.src.tar.gz"
	heimdal_cmd_src := exec.Command("curl", "-L", heimdal_src_url, "-o", "source.tar.gz")
	err = heimdal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	heimdal_cmd_direct := exec.Command("./binary")
	err = heimdal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: lmdb")
	exec.Command("latte", "install", "lmdb").Run()
	fmt.Println("Instalando dependencia: openldap")
	exec.Command("latte", "install", "openldap").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
