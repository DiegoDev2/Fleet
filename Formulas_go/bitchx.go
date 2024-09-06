package main

// Bitchx - Text-based, scriptable IRC client
// Homepage: https://bitchx.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBitchx() {
	// Método 1: Descargar y extraer .tar.gz
	bitchx_tar_url := "https://downloads.sourceforge.net/project/bitchx/ircii-pana/bitchx-1.2.1/bitchx-1.2.1.tar.gz"
	bitchx_cmd_tar := exec.Command("curl", "-L", bitchx_tar_url, "-o", "package.tar.gz")
	err := bitchx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bitchx_zip_url := "https://downloads.sourceforge.net/project/bitchx/ircii-pana/bitchx-1.2.1/bitchx-1.2.1.zip"
	bitchx_cmd_zip := exec.Command("curl", "-L", bitchx_zip_url, "-o", "package.zip")
	err = bitchx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bitchx_bin_url := "https://downloads.sourceforge.net/project/bitchx/ircii-pana/bitchx-1.2.1/bitchx-1.2.1.bin"
	bitchx_cmd_bin := exec.Command("curl", "-L", bitchx_bin_url, "-o", "binary.bin")
	err = bitchx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bitchx_src_url := "https://downloads.sourceforge.net/project/bitchx/ircii-pana/bitchx-1.2.1/bitchx-1.2.1.src.tar.gz"
	bitchx_cmd_src := exec.Command("curl", "-L", bitchx_src_url, "-o", "source.tar.gz")
	err = bitchx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bitchx_cmd_direct := exec.Command("./binary")
	err = bitchx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
