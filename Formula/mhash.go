package main

// Mhash - Uniform interface to a large number of hash algorithms
// Homepage: https://mhash.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMhash() {
	// Método 1: Descargar y extraer .tar.gz
	mhash_tar_url := "https://downloads.sourceforge.net/project/mhash/mhash/0.9.9.9/mhash-0.9.9.9.tar.gz"
	mhash_cmd_tar := exec.Command("curl", "-L", mhash_tar_url, "-o", "package.tar.gz")
	err := mhash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mhash_zip_url := "https://downloads.sourceforge.net/project/mhash/mhash/0.9.9.9/mhash-0.9.9.9.zip"
	mhash_cmd_zip := exec.Command("curl", "-L", mhash_zip_url, "-o", "package.zip")
	err = mhash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mhash_bin_url := "https://downloads.sourceforge.net/project/mhash/mhash/0.9.9.9/mhash-0.9.9.9.bin"
	mhash_cmd_bin := exec.Command("curl", "-L", mhash_bin_url, "-o", "binary.bin")
	err = mhash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mhash_src_url := "https://downloads.sourceforge.net/project/mhash/mhash/0.9.9.9/mhash-0.9.9.9.src.tar.gz"
	mhash_cmd_src := exec.Command("curl", "-L", mhash_src_url, "-o", "source.tar.gz")
	err = mhash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mhash_cmd_direct := exec.Command("./binary")
	err = mhash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
