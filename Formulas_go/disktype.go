package main

// Disktype - Detect content format of a disk or disk image
// Homepage: https://disktype.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDisktype() {
	// Método 1: Descargar y extraer .tar.gz
	disktype_tar_url := "https://downloads.sourceforge.net/project/disktype/disktype/9/disktype-9.tar.gz"
	disktype_cmd_tar := exec.Command("curl", "-L", disktype_tar_url, "-o", "package.tar.gz")
	err := disktype_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	disktype_zip_url := "https://downloads.sourceforge.net/project/disktype/disktype/9/disktype-9.zip"
	disktype_cmd_zip := exec.Command("curl", "-L", disktype_zip_url, "-o", "package.zip")
	err = disktype_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	disktype_bin_url := "https://downloads.sourceforge.net/project/disktype/disktype/9/disktype-9.bin"
	disktype_cmd_bin := exec.Command("curl", "-L", disktype_bin_url, "-o", "binary.bin")
	err = disktype_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	disktype_src_url := "https://downloads.sourceforge.net/project/disktype/disktype/9/disktype-9.src.tar.gz"
	disktype_cmd_src := exec.Command("curl", "-L", disktype_src_url, "-o", "source.tar.gz")
	err = disktype_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	disktype_cmd_direct := exec.Command("./binary")
	err = disktype_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
