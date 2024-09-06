package main

// Dos2unix - Convert text between DOS, UNIX, and Mac formats
// Homepage: https://waterlan.home.xs4all.nl/dos2unix.html

import (
	"fmt"
	
	"os/exec"
)

func installDos2unix() {
	// Método 1: Descargar y extraer .tar.gz
	dos2unix_tar_url := "https://waterlan.home.xs4all.nl/dos2unix/dos2unix-7.5.2.tar.gz"
	dos2unix_cmd_tar := exec.Command("curl", "-L", dos2unix_tar_url, "-o", "package.tar.gz")
	err := dos2unix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dos2unix_zip_url := "https://waterlan.home.xs4all.nl/dos2unix/dos2unix-7.5.2.zip"
	dos2unix_cmd_zip := exec.Command("curl", "-L", dos2unix_zip_url, "-o", "package.zip")
	err = dos2unix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dos2unix_bin_url := "https://waterlan.home.xs4all.nl/dos2unix/dos2unix-7.5.2.bin"
	dos2unix_cmd_bin := exec.Command("curl", "-L", dos2unix_bin_url, "-o", "binary.bin")
	err = dos2unix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dos2unix_src_url := "https://waterlan.home.xs4all.nl/dos2unix/dos2unix-7.5.2.src.tar.gz"
	dos2unix_cmd_src := exec.Command("curl", "-L", dos2unix_src_url, "-o", "source.tar.gz")
	err = dos2unix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dos2unix_cmd_direct := exec.Command("./binary")
	err = dos2unix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
