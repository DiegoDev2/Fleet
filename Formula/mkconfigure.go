package main

// MkConfigure - Lightweight replacement for GNU autotools
// Homepage: https://github.com/cheusov/mk-configure

import (
	"fmt"
	
	"os/exec"
)

func installMkConfigure() {
	// Método 1: Descargar y extraer .tar.gz
	mkconfigure_tar_url := "https://downloads.sourceforge.net/project/mk-configure/mk-configure/mk-configure-0.39.4/mk-configure-0.39.4.tar.gz"
	mkconfigure_cmd_tar := exec.Command("curl", "-L", mkconfigure_tar_url, "-o", "package.tar.gz")
	err := mkconfigure_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkconfigure_zip_url := "https://downloads.sourceforge.net/project/mk-configure/mk-configure/mk-configure-0.39.4/mk-configure-0.39.4.zip"
	mkconfigure_cmd_zip := exec.Command("curl", "-L", mkconfigure_zip_url, "-o", "package.zip")
	err = mkconfigure_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkconfigure_bin_url := "https://downloads.sourceforge.net/project/mk-configure/mk-configure/mk-configure-0.39.4/mk-configure-0.39.4.bin"
	mkconfigure_cmd_bin := exec.Command("curl", "-L", mkconfigure_bin_url, "-o", "binary.bin")
	err = mkconfigure_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkconfigure_src_url := "https://downloads.sourceforge.net/project/mk-configure/mk-configure/mk-configure-0.39.4/mk-configure-0.39.4.src.tar.gz"
	mkconfigure_cmd_src := exec.Command("curl", "-L", mkconfigure_src_url, "-o", "source.tar.gz")
	err = mkconfigure_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkconfigure_cmd_direct := exec.Command("./binary")
	err = mkconfigure_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bmake")
	exec.Command("latte", "install", "bmake").Run()
	fmt.Println("Instalando dependencia: makedepend")
	exec.Command("latte", "install", "makedepend").Run()
}
