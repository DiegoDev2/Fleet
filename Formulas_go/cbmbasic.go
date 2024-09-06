package main

// Cbmbasic - Commodore BASIC V2 as a scripting language
// Homepage: https://github.com/mist64/cbmbasic

import (
	"fmt"
	
	"os/exec"
)

func installCbmbasic() {
	// Método 1: Descargar y extraer .tar.gz
	cbmbasic_tar_url := "https://downloads.sourceforge.net/project/cbmbasic/cbmbasic/1.0/cbmbasic-1.0.tgz"
	cbmbasic_cmd_tar := exec.Command("curl", "-L", cbmbasic_tar_url, "-o", "package.tar.gz")
	err := cbmbasic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cbmbasic_zip_url := "https://downloads.sourceforge.net/project/cbmbasic/cbmbasic/1.0/cbmbasic-1.0.tgz"
	cbmbasic_cmd_zip := exec.Command("curl", "-L", cbmbasic_zip_url, "-o", "package.zip")
	err = cbmbasic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cbmbasic_bin_url := "https://downloads.sourceforge.net/project/cbmbasic/cbmbasic/1.0/cbmbasic-1.0.tgz"
	cbmbasic_cmd_bin := exec.Command("curl", "-L", cbmbasic_bin_url, "-o", "binary.bin")
	err = cbmbasic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cbmbasic_src_url := "https://downloads.sourceforge.net/project/cbmbasic/cbmbasic/1.0/cbmbasic-1.0.tgz"
	cbmbasic_cmd_src := exec.Command("curl", "-L", cbmbasic_src_url, "-o", "source.tar.gz")
	err = cbmbasic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cbmbasic_cmd_direct := exec.Command("./binary")
	err = cbmbasic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
