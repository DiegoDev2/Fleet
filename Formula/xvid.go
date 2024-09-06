package main

// Xvid - High-performance, high-quality MPEG-4 video library
// Homepage: https://labs.xvid.com/

import (
	"fmt"
	
	"os/exec"
)

func installXvid() {
	// Método 1: Descargar y extraer .tar.gz
	xvid_tar_url := "https://downloads.xvid.com/downloads/xvidcore-1.3.7.tar.bz2"
	xvid_cmd_tar := exec.Command("curl", "-L", xvid_tar_url, "-o", "package.tar.gz")
	err := xvid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xvid_zip_url := "https://downloads.xvid.com/downloads/xvidcore-1.3.7.tar.bz2"
	xvid_cmd_zip := exec.Command("curl", "-L", xvid_zip_url, "-o", "package.zip")
	err = xvid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xvid_bin_url := "https://downloads.xvid.com/downloads/xvidcore-1.3.7.tar.bz2"
	xvid_cmd_bin := exec.Command("curl", "-L", xvid_bin_url, "-o", "binary.bin")
	err = xvid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xvid_src_url := "https://downloads.xvid.com/downloads/xvidcore-1.3.7.tar.bz2"
	xvid_cmd_src := exec.Command("curl", "-L", xvid_src_url, "-o", "source.tar.gz")
	err = xvid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xvid_cmd_direct := exec.Command("./binary")
	err = xvid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
