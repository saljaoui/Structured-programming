import { places } from "./where-do-we-go.data.js";
const sections = [];
let nameElement = null;
export const explore = () => {

  const sortedPlaces = sort()
  console.log(sections);
  
  addEventListener("scroll", (event) => {
    const scrollY = window.scrollY;

    // Check each section to see if it's at the end
    sections.forEach((section, index) => {
        const sectionTop = section.offsetTop;
        const sectionHeight = section.offsetHeight;

        // Check if we are at the bottom of the current section
        if (scrollY >= sectionTop + (sectionHeight/2) - window.innerHeight) {
          if (nameElement) {
            nameElement.remove();
          }
            //  console.log(`Reached the end of section ${index + 1}: ${section.textContent}`);
               nameElement = document.createElement('a')
  nameElement.className = "location"
  nameElement.textContent = places[index].name
  document.body.append(nameElement)
        }
      });
  })

  // const nameElement = document.createElement('a')
  // nameElement.className = "location"
  // nameElement.textContent = "soufian"
  // document.body.append(nameElement)
    
  sortedPlaces.forEach(place => {
      const section = document.createElement('section');
      const firstPart = place.name.split(",")[0].trim();
      section.style.backgroundImage = `url('${firstPart.toLowerCase().replace(/\s+/g, '-')}.jpg')`
;
      section.style.backgroundSize = 'cover';
      section.style.backgroundPosition = 'center';
      section.className = "section"
      section.style.display = 'flex';
      section.style.justifyContent = 'center';
      section.style.alignItems = 'center';
      sections.push(section);
      document.body.appendChild(section);
    });

      
    
  };

  const sort = () => {
    const northPlaces = places.filter(place => place.coordinates.includes("N"))
    const southPlaces = places.filter(place => place.coordinates.includes("S"))
    northPlaces.sort((a, b) => {
        if (a.coordinates > b.coordinates) return -1
        if (b.coordinates > a.coordinates) return 1   
    })
    southPlaces.sort((b, a) => {
        if (a.coordinates > b.coordinates) return -1
        if (b.coordinates > a.coordinates) return 1   
    })
    return northPlaces.concat(southPlaces)
}
