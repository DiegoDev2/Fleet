package main

// Lasi - C++ stream output interface for creating Postscript documents
// Homepage: https://www.unifont.org/lasi/

import (
	"fmt"
	
	"os/exec"
)

func installLasi() {
	// Método 1: Descargar y extraer .tar.gz
	lasi_tar_url := "https://downloads.sourceforge.net/project/lasi/lasi/1.1.3%20Source/libLASi-1.1.3.tar.gz"
	lasi_cmd_tar := exec.Command("curl", "-L", lasi_tar_url, "-o", "package.tar.gz")
	err := lasi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lasi_zip_url := "https://downloads.sourceforge.net/project/lasi/lasi/1.1.3%20Source/libLASi-1.1.3.zip"
	lasi_cmd_zip := exec.Command("curl", "-L", lasi_zip_url, "-o", "package.zip")
	err = lasi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lasi_bin_url := "https://downloads.sourceforge.net/project/lasi/lasi/1.1.3%20Source/libLASi-1.1.3.bin"
	lasi_cmd_bin := exec.Command("curl", "-L", lasi_bin_url, "-o", "binary.bin")
	err = lasi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lasi_src_url := "https://downloads.sourceforge.net/project/lasi/lasi/1.1.3%20Source/libLASi-1.1.3.src.tar.gz"
	lasi_cmd_src := exec.Command("curl", "-L", lasi_src_url, "-o", "source.tar.gz")
	err = lasi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lasi_cmd_direct := exec.Command("./binary")
	err = lasi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
