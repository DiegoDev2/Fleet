package main

// Croaring - Roaring bitmaps in C (and C++)
// Homepage: https://roaringbitmap.org

import (
	"fmt"
	
	"os/exec"
)

func installCroaring() {
	// Método 1: Descargar y extraer .tar.gz
	croaring_tar_url := "https://github.com/RoaringBitmap/CRoaring/archive/refs/tags/v4.1.2.tar.gz"
	croaring_cmd_tar := exec.Command("curl", "-L", croaring_tar_url, "-o", "package.tar.gz")
	err := croaring_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	croaring_zip_url := "https://github.com/RoaringBitmap/CRoaring/archive/refs/tags/v4.1.2.zip"
	croaring_cmd_zip := exec.Command("curl", "-L", croaring_zip_url, "-o", "package.zip")
	err = croaring_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	croaring_bin_url := "https://github.com/RoaringBitmap/CRoaring/archive/refs/tags/v4.1.2.bin"
	croaring_cmd_bin := exec.Command("curl", "-L", croaring_bin_url, "-o", "binary.bin")
	err = croaring_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	croaring_src_url := "https://github.com/RoaringBitmap/CRoaring/archive/refs/tags/v4.1.2.src.tar.gz"
	croaring_cmd_src := exec.Command("curl", "-L", croaring_src_url, "-o", "source.tar.gz")
	err = croaring_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	croaring_cmd_direct := exec.Command("./binary")
	err = croaring_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
