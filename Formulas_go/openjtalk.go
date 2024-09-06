package main

// OpenJtalk - Japanese text-to-speech system
// Homepage: https://open-jtalk.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installOpenJtalk() {
	// Método 1: Descargar y extraer .tar.gz
	openjtalk_tar_url := "https://downloads.sourceforge.net/project/open-jtalk/Open%20JTalk/open_jtalk-1.11/open_jtalk-1.11.tar.gz"
	openjtalk_cmd_tar := exec.Command("curl", "-L", openjtalk_tar_url, "-o", "package.tar.gz")
	err := openjtalk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openjtalk_zip_url := "https://downloads.sourceforge.net/project/open-jtalk/Open%20JTalk/open_jtalk-1.11/open_jtalk-1.11.zip"
	openjtalk_cmd_zip := exec.Command("curl", "-L", openjtalk_zip_url, "-o", "package.zip")
	err = openjtalk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openjtalk_bin_url := "https://downloads.sourceforge.net/project/open-jtalk/Open%20JTalk/open_jtalk-1.11/open_jtalk-1.11.bin"
	openjtalk_cmd_bin := exec.Command("curl", "-L", openjtalk_bin_url, "-o", "binary.bin")
	err = openjtalk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openjtalk_src_url := "https://downloads.sourceforge.net/project/open-jtalk/Open%20JTalk/open_jtalk-1.11/open_jtalk-1.11.src.tar.gz"
	openjtalk_cmd_src := exec.Command("curl", "-L", openjtalk_src_url, "-o", "source.tar.gz")
	err = openjtalk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openjtalk_cmd_direct := exec.Command("./binary")
	err = openjtalk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
