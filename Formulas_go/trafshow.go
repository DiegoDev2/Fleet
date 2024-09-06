package main

// Trafshow - Continuous network traffic display
// Homepage: https://web.archive.org/web/20130707021442/soft.risp.ru/trafshow/index_en.shtml

import (
	"fmt"
	
	"os/exec"
)

func installTrafshow() {
	// Método 1: Descargar y extraer .tar.gz
	trafshow_tar_url := "https://pkg.freebsd.org/ports-distfiles/trafshow-5.2.3.tgz"
	trafshow_cmd_tar := exec.Command("curl", "-L", trafshow_tar_url, "-o", "package.tar.gz")
	err := trafshow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trafshow_zip_url := "https://pkg.freebsd.org/ports-distfiles/trafshow-5.2.3.tgz"
	trafshow_cmd_zip := exec.Command("curl", "-L", trafshow_zip_url, "-o", "package.zip")
	err = trafshow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trafshow_bin_url := "https://pkg.freebsd.org/ports-distfiles/trafshow-5.2.3.tgz"
	trafshow_cmd_bin := exec.Command("curl", "-L", trafshow_bin_url, "-o", "binary.bin")
	err = trafshow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trafshow_src_url := "https://pkg.freebsd.org/ports-distfiles/trafshow-5.2.3.tgz"
	trafshow_cmd_src := exec.Command("curl", "-L", trafshow_src_url, "-o", "source.tar.gz")
	err = trafshow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trafshow_cmd_direct := exec.Command("./binary")
	err = trafshow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
