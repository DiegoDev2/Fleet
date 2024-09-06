package main

// Seexpr - Embeddable expression evaluation engine
// Homepage: https://wdas.github.io/SeExpr/

import (
	"fmt"
	
	"os/exec"
)

func installSeexpr() {
	// Método 1: Descargar y extraer .tar.gz
	seexpr_tar_url := "https://github.com/wdas/SeExpr/archive/refs/tags/v3.0.1.tar.gz"
	seexpr_cmd_tar := exec.Command("curl", "-L", seexpr_tar_url, "-o", "package.tar.gz")
	err := seexpr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	seexpr_zip_url := "https://github.com/wdas/SeExpr/archive/refs/tags/v3.0.1.zip"
	seexpr_cmd_zip := exec.Command("curl", "-L", seexpr_zip_url, "-o", "package.zip")
	err = seexpr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	seexpr_bin_url := "https://github.com/wdas/SeExpr/archive/refs/tags/v3.0.1.bin"
	seexpr_cmd_bin := exec.Command("curl", "-L", seexpr_bin_url, "-o", "binary.bin")
	err = seexpr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	seexpr_src_url := "https://github.com/wdas/SeExpr/archive/refs/tags/v3.0.1.src.tar.gz"
	seexpr_cmd_src := exec.Command("curl", "-L", seexpr_src_url, "-o", "source.tar.gz")
	err = seexpr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	seexpr_cmd_direct := exec.Command("./binary")
	err = seexpr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
