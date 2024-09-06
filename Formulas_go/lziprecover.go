package main

// Lziprecover - Data recovery tool and decompressor for files in the lzip compressed data format
// Homepage: https://www.nongnu.org/lzip/lziprecover.html

import (
	"fmt"
	
	"os/exec"
)

func installLziprecover() {
	// Método 1: Descargar y extraer .tar.gz
	lziprecover_tar_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lziprecover/lziprecover-1.24.tar.gz"
	lziprecover_cmd_tar := exec.Command("curl", "-L", lziprecover_tar_url, "-o", "package.tar.gz")
	err := lziprecover_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lziprecover_zip_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lziprecover/lziprecover-1.24.zip"
	lziprecover_cmd_zip := exec.Command("curl", "-L", lziprecover_zip_url, "-o", "package.zip")
	err = lziprecover_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lziprecover_bin_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lziprecover/lziprecover-1.24.bin"
	lziprecover_cmd_bin := exec.Command("curl", "-L", lziprecover_bin_url, "-o", "binary.bin")
	err = lziprecover_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lziprecover_src_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lziprecover/lziprecover-1.24.src.tar.gz"
	lziprecover_cmd_src := exec.Command("curl", "-L", lziprecover_src_url, "-o", "source.tar.gz")
	err = lziprecover_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lziprecover_cmd_direct := exec.Command("./binary")
	err = lziprecover_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lzip")
exec.Command("latte", "install", "lzip")
}
